# 简介
- **功能**

    - 语音服务（文本转语音TTS）
    - 短信服务
    - 邮件服务
    
- **核心软件**

    Python3.7.3, Flask1.1.1, Gunicorn20.0.4, Gevent1.4.0，[阿里云开源SDK](https://github.com/aliyun/aliyun-openapi-python-sdk)
    
# 调试
* **调试运行**

  ```bash
  cd $HOME/aliyun-sms/;git pull;python main.py
  或
  cd $HOME/aliyun-sms/;git pull;nohup python main.py > out.log 2>&1 &
  ```

* **命令行调试**

  ```bash
  export FLASK_APP=main（默认是app模块的）
  $ flask run
  $ flask shell
  ```

* **环境切换**

  ```bash
  app = create_app(os.getenv('FLASK_CONFIG_ENV') or 'dev') #pro
  ```

* **正式运行**

  ```bash
  cd $HOME/aliyun-sms/;git pull;gunicorn main:app -c gunicorn.conf.py >> /data/logs/aliyun-sms/gunicorn.log 2>&1 &
  ```

  

# 接口说明

* ## **冰鉴短信服务接口**

  ```bash
  https://notify.icekredit.com/sms（生产环境）
  http://172.16.21.100:8010/sms（开发环境）
  ```
  
  Python调用参考代码如下：
  
    ```python
    import requests
    
    def sendSMS(GroupID, PhoneNumbers ,  SignName , TemplateCode, TemplateParam):
        payload = { "GroupID": GroupID, "PhoneNumbers": PhoneNumbers , "SignName" : SignName , "TemplateCode":TemplateCode, "TemplateParam" : TemplateParam  }
        try:
            ret = requests.post("https://notify.icekredit.com/sms", json=payload, headers={'Content-Type': 'application/json'})
            print(ret.text)
        except Exception as e:
            print("发送失败:%s" %(e))
    
    if __name__ == '__main__':
        #富民银行短信模板:SMS_184105379
      sendSMS("bc-fumin", 18683441008, "冰鉴科技", "SMS_184105379", "654321")
    ```
  
  **参数说明**  [请参考阿里云SendSmsAPI文档](https://help.aliyun.com/document_detail/101414.html?spm=a2c4g.11186623.6.624.2f8b152cTdaXr4)
  
  **GroupID**：*字符串*，授权的组ID，只有在接口配置过该组ID才能调用，一般用来判断调用来源**PhoneNumbers**：*String*，必填，接收短信的手机号码
  
  **SignName**： *String*，必填，短信签名名称。请在控制台**签名管理**页面**签名名称**一列查看
  
  **TemplateCode**： *String*，必填，短信模板ID。请在控制台**模板管理**页面**模板CODE**一列查看
  
  **TemplateParam**：*String*，必填，短信模板变量对应的实际值，这里填写验证码即可
  
  **返回值:**
  
  ```json
  成功:
      {"Message":"OK","RequestId":"335609C6-811E-4BA3-B1C0-F34C9FB98B7D","BizId":"957805682295447472^0","Code":"OK"}
  失败:
      {"Message":"请联系管理员确认GroupID是否正确","RequestId":"null","BizId": "null","Code": "ERROR"}
  ```
  
  **Message**：状态码的描述
  
  **RequestId**：阿里云产生的请求ID，自定义错误该值为空
  
  **BizId**：发送回执ID，可根据该ID在接口QuerySendDetails中查询具体的发送状态，自定义错误该值为空
  
  **Code**：请求状态码。
  
  * 返回OK代表请求成功。
  * 其他错误码详见[错误码列表](https://help.aliyun.com/document_detail/101346.html)。
  
* ## **冰鉴语音服务接口（文本转语音TTS）**

  ```bash
  https://notify.icekredit.com/tts（生产环境）
  http://172.16.21.100:8010/tts（开发环境）
  ```

  Python调用参考代码如下：

    ```python
  import requests
  
  def sendSMS(GroupID,  CalledNumber , TtsCode, TtsParam):
      payload = { "GroupID": GroupID, "CalledNumber": CalledNumber , "TtsCode" : TtsCode , "TtsParam":TtsParam }
      try:
          ret = requests.post("https://notify.icekredit.com/tts", json=payload, headers={'Content-Type': 'application/json'})
          print(ret.text)
      except Exception as e:
          print("发送失败:%s" %(e))
  
  if __name__ == '__main__':
      sendSMS("bc-fmb", 18683441008, "TTS_134311476", {"errorcode": "TTS测试", 'count': "10"})
    ```

  **参数说明**  [请参考阿里云SingleCallByTts文档](https://help.aliyun.com/document_detail/114035.html)

  请求参数：

  | 名称             | 类型   | 是否必选 | 示例值              | 描述                                                         |
  | :--------------- | :----- | :------- | :------------------ | :----------------------------------------------------------- |
  | **GroupID**      | String | 是       | bc-fmb              | 自定义参数, 授权的组ID，只有在接口配置过该组ID才能调用，一般用来判断调用来源 |
  | **CalledNumber** | String | 是       | 13700000000         | 被叫号码。仅支持中国大陆号码。                               |
  | **TtsCode**      | String | 是       | TTS_10001           | 已审核通过的语音验证码模板ID，您可以在[文本转语音模板页面](https://dyvms.console.aliyun.com/dyvms.htm#/template)查看模板ID。 |
  | **TtsParam**     | String | 否       | {“AckNum”:”123456”} | 模板中的变量参数，格式为JSON。                               |

  返回数据：

  | 名称          | 类型   | 示例值                               | 描述                                                         |
  | :------------ | :----- | :----------------------------------- | :----------------------------------------------------------- |
  | **RequestId** | String | D9CB3933-9FE3-4870-BA8E-2BEE91B69D23 | 请求ID。(自定义错误该值为空)                                 |
  | **CallId**    | String | 116012354148^1028137841xx            | 此次通话的唯一回执ID，可以用此ID通过接口**QueryCallDetailByCallId**查询呼叫详情。（自定义错误该值为空） |
  | **Code**      | String | OK                                   | 请求状态码。返回OK代表请求成功。其他错误码详见[错误码列表](https://help.aliyun.com/document_detail/112502.html)。（自定义错误一般为ERROR） |
  | **Message**   | String | OK                                   | 状态码的描述。                                               |
  
  
  
* ## **冰鉴邮件接口**

  ```bash
  https://notify.icekredit.com/mail（生产环境）
  http://172.16.21.100:8010/mail（开发环境）
  ```

  Python调用参考代码如下：

    ```python
  import requests
  
  def sendMail(group_id, subject,  to , cc, body):
      payload = { "group_id": group_id, "subject": subject, "to": to , "cc" : cc , "body":body }
      try:
          ret = requests.post("http://172.16.21.100:8010/mail", json=payload, headers={'Content-Type': 'application/json'})
          print(ret.text)
      except Exception as e:
          print("发送失败:%s" %(e))
  
  if __name__ == '__main__':
      sendMail("bc-fmb","邮件测试MailTest", ["song_xiaobing@icekredit.com", "yan_hao@icekredit.com"], ["song_xiaobing@icekredit.com"], "邮件正文MailBody")
    ```

  **参数说明**

  请求参数：

  | 名称         | 类型   | 是否必选 | 示例值                                  | 描述                                                         |
  | ------------ | ------ | -------- | --------------------------------------- | ------------------------------------------------------------ |
  | **group_id** | String | 是       | bc-fmb                                  | 自定义参数, 授权的组ID，只有在接口配置过该组ID才能调用，一般用来判断调用来源 |
  | **subject**  | String | 是       | 邮件主题测试                            | 邮件主题                                                     |
  | **to**       | List   | 是       | ["a@icekredit.com", "b@icekredit.com,"] | 邮件接收人列表                                               |
  | **cc**       | List   | 否       | ["a@icekredit.com", "b@icekredit.com,"] | 邮件抄送人列表                                               |
  | **body**     | String | 是       | 正文测试                                | 邮件正文                                                     |

  返回数据：

    ```json
    {
      "Code": "OK", 
      "Message": "邮件已发送", 
      "ThreadIdentifier": "139691417802136", 
      "ThreadName": "Thread-2"
    }
    ```

  参数详解：

  | 名称                 | 类型   | 示例值          | 描述                                                         |
  | :------------------- | :----- | :-------------- | ------------------------------------------------------------ |
  | **ThreadIdentifier** | String | 139691417802136 | 后台邮件线程标识号，如果未正常收到邮件，可以将此ID发给管理员确认详细原因 |
  | **ThreadName**       | String | Thread-2        | 后台邮件线程名                                               |
  | **Code**             | String | OK              | 请求状态码，返回OK代表请求成功，其他表示异常                 |
  | **Message**          | String | OK              | 状态码的描述                                                 |