# API_HDU_Notice

对 `Tag` 的定义，既本接口所称呼的频道，频道下有子频道，对应不同站点及其栏目

订阅父频道会接收到所有子频道的推送

## 接口索引

[GET /tags 查看所有频道的tag](#get-tags-获取所有的通知频道)

`/sub` 订阅通知

- [GET /sub 查看订阅信息](#get-sub-查看订阅信息)
- [POST /sub 订阅通知](#post-sub-订阅通知)
- [DELETE /sub 取消订阅](#delete-sub-取消订阅)

[GET /channel 查看频道通知](#get-channel-查看频道通知)

[GET /feed 查看频道通知](#get-feed-查看频道通知)

### GET `/tags` 获取所有的通知频道

Response(200):

```json
{
  "data": {
    "Type": "学校",
    "Site": "杭州电子科技大学",
    "Section": "杭电",
    "Tag": "",
    "Items": [
      {
        "Type": "学院",
        "Site": "杭电浙江磁性材料研究院",
        "Section": "",
        "Tag": "icam",
        "Items": [
          {
            "Type": "学院",
            "Site": "学术会议",
            "Section": "学术会议",
            "Tag": "icam_scholar"
          }
        ]
      }
    ]
  },
  "error": 0,
  "msg": "success"
}
```

## `/sub` 订阅通知

### GET `/sub` 查看订阅信息

Response(200):

```json
{
  "data": [
    {
      "StaffID": "19011141",
      "Tag": "jsj",
      "Name": "计算机学院总版"
    }
  ],
  "error": 0,
  "msg": "success"
}
```

### POST `/sub` 订阅通知

订阅父频道后会删除对子频道的订阅（避免重复推送）

如: 订阅 `jsj` Tag 后，会删除对 `jsj_notice` `jsj_news` Tag 的订阅

Query

- `tag` 通知频道Tag

Response(200):

```json
{
  "data": null,
  "error": 0,
  "msg": "success"
}
```

### DELETE `/sub` 取消订阅

删除父频道的订阅会删除父频道本身及其子频道的订阅

Query

- `tag` 通知频道Tag

Response(200):

```json
{
  "data": null,
  "error": 0,
  "msg": "success"
}
```

### GET `/channel` 查看频道通知

Query

- `tag` 通知频道Tag
- `page` 页数 默认第1页 选填
- `rpp` 每页条数 默认25 选填

Response(200):

```json
{
  "data": {
    "Tag": "jsj",
    "Title": "计算机学院",
    "URL": "https://computer.hdu.edu.cn",
    "data": {
      "PageNow": 1,
      //当前页面
      "PageCount": 7,
      //页面总数
      "RawCount": 157,
      //信息条数
      "RawPerPage": 25,
      //每页条数
      "ResultSet": [
        {
          "Content": "关于2019年个人所得税综合所得年度汇算温馨提醒：   2019年个人所得税综合所得年度汇算已开始，现将有关情况提醒如下：1、汇算时间     汇算时间是2020.3.1-6.30。受新冠疫情的影响，浙江省2020.4.1开始启动个人手机APP网上申报。    注：6月30日前完成提交，否则影响个人诚信。2、申报范围   全体在编教职工、退休返聘人员、学生（博士生、硕士生、本科生），在校内和校外获得的工资薪金、劳务报酬、稿酬、特许权使用费等四项所得。   注：外籍人员在国内低于180天，不作汇算；高于180天，要作汇算。凭护照到税务大厅取得注册码--下载汇算APP。3、相关的税收政策    年收入低于6万元，可在个人所得税按综合所得年度汇算APP上进行简易申报（5月底关闭简易申报系统）；年收入高于6万元，采用标准申报。   全年一次性年终奖金可选择“单独”计税方式或“全部并入综合所得”计税方式，测算两种方式的扣税额。一般人员请按“单独”计税方式可能扣税较低。   有捐赠的可以抵扣税额。                                            计算机学院                                        2020年5月15日",
          "Url": "https://computer.hdu.edu.cn/2020/0515/c1377a107715/page.htm",
          //信息源地址
          "Date": "2020-05-15T08:00:00+08:00",
          //信息时间
          "Site": "https://computer.hdu.edu.cn/1377/list.htm",
          //信息所属网站地址
          "Tag": "jsj_notice",
          "Title": "关于2019年个人所得税综合所得年度汇算温馨提醒",
          //信息标题
          "SiteName": "通知公告"
          //信息所属网站板块
        }
      ],
      "FirstPage": true,
      //是否第一页
      "LastPage": false,
      //是否最后一页
      "Empty": false,
      //是否空页面
      "StartRow": 25,
      //开始条数
      "EndRow": 49
      //结束条数
    }
  },
  "error": 0,
  "msg": "success"
}
```

### GET `/feed` 查看频道通知

适用于 Atom Feed

Query

- `tag` 通知频道Tag

Response(200):

```xml

<feed xmlns="https://www.w3.org/2005/Atom">
    <title>计算机学院</title>
    <id></id>
    <link href="https://computer.hdu.edu.cn"></link>
    <updated>2020-05-24T02:27:25+08:00</updated>
    <entry>
        <title>关于2019年个人所得税综合所得年度汇算温馨提醒</title>
        <id>jsj_notice-1782</id>
        <link href="https://computer.hdu.edu.cn/2020/0515/c1377a107715/page.htm"></link>
        <published>2020-05-15T08:00:00+08:00</published>
        <updated>2020-05-15T17:31:45+08:00</updated>
        <content type="">关于2019年个人所得税综合所得年度汇算温馨提醒：   2019年个人所得税综合所得年度汇算已开始，现将有关情况提醒如下：1、汇算时间    
            汇算时间是2020.3.1-6.30。受新冠疫情的影响，浙江省2020.4.1开始启动个人手机APP网上申报。   
            注：6月30日前完成提交，否则影响个人诚信。2、申报范围  
            全体在编教职工、退休返聘人员、学生（博士生、硕士生、本科生），在校内和校外获得的工资薪金、劳务报酬、稿酬、特许权使用费等四项所得。  
            注：外籍人员在国内低于180天，不作汇算；高于180天，要作汇算。凭护照到税务大厅取得注册码--下载汇算APP。3、相关的税收政策   
            年收入低于6万元，可在个人所得税按综合所得年度汇算APP上进行简易申报（5月底关闭简易申报系统）；年收入高于6万元，采用标准申报。  
            全年一次性年终奖金可选择“单独”计税方式或“全部并入综合所得”计税方式，测算两种方式的扣税额。一般人员请按“单独”计税方式可能扣税较低。  
            有捐赠的可以抵扣税额。                                            计算机学院                                 
                  2020年5月15日
        </content>
    </entry>
</feed>
```

错误对照:

| HTTPStatus | Error | Describe      |
|------------|-------|---------------|
| 404        | 40400 | 请求参数 `tag` 错误 |
| 403        | 40301 | 无效请求          |