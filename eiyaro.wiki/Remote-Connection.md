## Remote Connection

if you build a full-node in a server and want to call the API remotely, you must attach an access-token to validate. or you will get EY860 error(Request could not be authenticated)

how to get access-token? you can use SSH to server and call the API `create-access-token`  locally to create access-token,the way is below

`./eiyarocli create-access-token test` 

or

`curl -X POST create-access-token -d '{"id":"test"}'`

the response is like below

```
  "created_at": "2018-05-18T16:00:25.284677605+08:00",
  "id": "test",
  "token":"test:fe50927ddaa5bcca77021e9f50fa5ef236a6140c012d1fe2eb9241f61a9228e4
```
username : `test` 
password  : `fe50927ddaa5bcca77021e9f50fa5ef236a6140c012d1fe2eb9241f61a9228e4`

now you have got the access-token, if you use a tool(like postman) to send the request,you can set Authorization as Basic Auth, and fill the username and password.

if you want use code to request , you can refer to the code below(Java version):
```
String auth = Username + ":" + Password;
byte[] encodedAuth = Base64.encodeBase64(auth.getBytes(Charset.forName("US-ASCII")));
String authHeader = "Basic " + new String(encodedAuth);
Map<String, String> header = new LinkedHashMap<String, String>();
header.put("Authorization", authHeader);
```

**Notice : remember add the string "Basic" to request header**



