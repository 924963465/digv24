package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv24/global"
	"github.com/liuhongdi/digv24/pkg/image"
	"github.com/liuhongdi/digv24/pkg/result"
	"github.com/liuhongdi/digv24/pkg/validCheck"
	"github.com/liuhongdi/digv24/request"
	"strconv"
)

type ImageController struct{}

func NewImageController() ImageController {
	return ImageController{}
}
//上传单张图片
func (a *ImageController) UploadOne(c *gin.Context) {
	resultRes := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, errs := validCheck.BindAndValid(c, &param)
	if !valid {
		resultRes.Error(400,errs.Error())
		return
	}

    //save image
	f, err := c.FormFile("f1s")
	//错误处理
	if err != nil {
		fmt.Println(err.Error())
		resultRes.Error(1,"图片上传失败")
		} else {
             //将文件保存至本项目根目录中
			  idstr:=strconv.FormatUint(param.ID, 10)
			  destImage := global.ArticleImageSetting.UploadDir+"/"+idstr+".jpg"
              err := c.SaveUploadedFile(f, destImage)
              if (err != nil){
              	  fmt.Println("save err:")
				  fmt.Println(err)
				  resultRes.Error(1,"图片保存失败")
			  } else {
			  	  //make tmb
			  	  orig:= destImage
			  	  dest := global.ArticleImageSetting.TmbDir+"/"+idstr+".jpg"
				  err := image.ConvertByLong(orig,dest,300)
				  if (err != nil){
				  	  fmt.Println(err)
				  }
			  	  origUrl := global.ArticleImageSetting.ImageHost+"/static/ware/orig/"+idstr+".jpg"
				  tmbUrl := global.ArticleImageSetting.ImageHost+"/static/ware/tmb/"+idstr+".jpg"
				  resultRes.Success(gin.H{"origurl":origUrl,"tmburl":tmbUrl})
			  }
      }
	return
}
