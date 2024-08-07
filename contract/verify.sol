    // SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Verification{

        address owner;
        constructor(){
            owner=msg.sender;
        }

        enum DocStatus{
            accepted,
            rejected,
            pending
        }
        struct User{
            string publicKey;
        }

        struct Institution{
            address publicAddr;
            string  publicKey;
            string  name;
            bool    approved;
        }
        struct Document{
            address     requester;
            address     verifiedBy;
            string      institution; 
            string      name;
            string      description; 
            string      encrpytedIPFSHash;
            DocStatus   status;
            uint        index;
        }

        //each institution can only have one verifier, at least for now. 
        mapping(string=>Institution) public institutions;
        mapping(address=>User) private users;
        mapping(string=>uint) private documentList;
        mapping(address=>bool) userList;
        //jack-ass-hack
        Document[]  allDocuments;
        address[]   requesters;
        address[]   verifiedBy; 
        string[]    names;
        string[]    institution;
        string[]    descriptions;
        string[]    encrpytedIPFSHashes; 
        DocStatus[] status;

        uint docIndexCounter=0;

        function registerAsUser(string calldata _publicKey) public{
            users[msg.sender]=User({
                publicKey:  _publicKey
            });
            userList[msg.sender]=true;
        }
        
        function registerInstitution(string memory _publicKey, string memory _name) public{
            institutions[_name]=Institution({
                publicAddr: msg.sender,
                publicKey:  _publicKey,
                name:       _name,
                approved:   false
            });
        }


        function getInstituePublicKey(string memory _name) public view returns(string memory pubKey){
            return institutions[_name].publicKey;
        }

        function getUserPublicKey(address userAddr)public view returns (string memory){
            return users[userAddr].publicKey;        
        }

        function approveVerifier(string memory _name)public{
            require(msg.sender==owner,"Only admin can perfom this action");
            institutions[_name].approved=true;
        }

        function addDocument(string memory _EncryptedIPFSHash, string memory _institute,string memory _name ,string memory _description) public{
            bool isUser=userList[msg.sender];
            require(isUser==true,"register first to verify");
            documentList[_EncryptedIPFSHash]=docIndexCounter;

            requesters.push(msg.sender);
            verifiedBy.push(address(0));
            institution.push(_institute);
            names.push(_name);
            descriptions.push(_description);
            status.push(DocStatus.pending);
            encrpytedIPFSHashes.push(_EncryptedIPFSHash);
            docIndexCounter++;
        }
        //returns all the documents
        function getDocuments()public view returns(address[] memory requester ,address[] memory verifer ,string[] memory institute,string[] memory ipfs,string[] memory name,string[] memory desc,DocStatus[] memory stats){
            return (requesters,verifiedBy,institution,encrpytedIPFSHashes,names,descriptions,status);
        }
        
        function verifyDocument(string memory _institute, string memory _ipfs, DocStatus _status)public payable {
            require(institutions[_institute].approved==true && institutions[_institute].publicAddr==msg.sender);
            uint index=documentList[_ipfs];
            status[index]=_status;
            verifiedBy[index]=msg.sender;
        }

        function isApprovedInstitute(string memory name)public view returns (bool){
            return institutions[name].approved==true;
        }

}        
