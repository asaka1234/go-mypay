整体流程
==============
1. server负责对请求参数做sign签名. 随后发请求给psp
4. server收到回调


鉴权
==============
1. json body的md5签名


回调地址
==============
在psp端指定的callback地址, 所以无法在pre-order中动态指定


Comment
===============
1. both support deposit && withdrawl
2. 只支持application/json格式.