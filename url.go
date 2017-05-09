package wechat

var urlMap = map[int64]map[string]string{
    QY: {
        "GetToken":"https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
        "GetServerList":"https://qyapi.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s",
    },
    MP: {
        "GetToken":"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
        "GetServerList":"https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=%s",
        "UploadMedia":"https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s",             //新增临时素材,POST HTTPS
        "GetMedia":"https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s",               //获取临时素材,GET HTTPS
        "UploadMaterial":"https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s",             //新增永久素材,POST HTTPS
        "GetMaterial":"https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s",            //获取永久素材,POST HTTPS
        "DelMaterial":"https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s",            //删除永久素材,POST HTTPS
        "UpdateNews":"https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=%s",              //修改永久图文素材, POST HTTPS
        "GetMaterialCount":"https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s",  //获取素材总数, GET HTTPS
        "GetMaterialList":"https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s",   //获取素材列表

    },
    SR: {
    },
}


type UrlMap struct {
    GetTokenUrl            string
    GetServerListUrl       string
    UploadMediaUrl         string
    GetMediaUrl            string
    UploadMaterialUrl      string
    GetMaterialUrl         string
    DelMaterialUrl         string
    UpdateNewsUrl          string
    GetMaterialCountUrl    string
    GetMaterialListUrl     string
}




func Url(_type int64,name string) string {
    return urlMap[_type][name]
}
