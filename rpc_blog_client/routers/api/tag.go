package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"rpc_blog_client/models"
	"rpc_blog_client/pkg/app"
	"rpc_blog_client/pkg/e"
	"rpc_blog_client/pkg/util"
	"rpc_blog_client/rpc"
)

func GetTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID不合法")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusOK, app.GetResponse(e.INVALID_PARAMS, nil))
		return
	}

	tag, code := rpc.GetTag(id)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, app.GetResponse(code, nil))
		return
	}
	c.JSON(http.StatusOK, app.GetResponse(code, tag))
}

type GetTagForm struct {
	Name       string `form:"name" valid:"MaxSize(255)"`
	State      int    `form:"state" valid:"Range(0,1)"`
	ModifiedBy string `form:"modified_by" valid:"MaxSize(100)"`
	CreatedBy  string `form:"created_by" valid:"MaxSize(100)"`
}

func GetTags(c *gin.Context) {
	var tagForm GetTagForm
	httpCode, code := app.BindAndValid(c, &tagForm)
	if code != e.SUCCESS {
		c.JSON(httpCode, app.GetResponse(code, nil))
		return
	}

	maps := make(map[string]interface{})
	if tagForm.Name != "" {
		maps["name"] = tagForm.Name
	}
	if tagForm.State != 0 {
		maps["state"] = tagForm.State
	}
	if tagForm.ModifiedBy != "" {
		maps["modified_by"] = tagForm.ModifiedBy
	}
	if tagForm.CreatedBy != "" {
		maps["created_by"] = tagForm.CreatedBy
	}
	offset, limit := util.GetPage(c)
	data, code := rpc.GetTags(offset, limit, maps)
	c.JSON(http.StatusOK, app.GetResponse(code, data))
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(255)"`
	State     int    `form:"state" valid:"Range(0,1)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
}

// @Summary 新增文章标签
// @Produce json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query string false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /tag [post]
func AddTag(c *gin.Context) {
	var tagForm AddTagForm
	httpCode, code := app.BindAndValid(c, &tagForm)
	if code != e.SUCCESS {
		c.JSON(httpCode, app.GetResponse(code, nil))
		return
	}

	tag := models.Tag{
		Name:      tagForm.Name,
		CreatedBy: tagForm.CreatedBy,
		State:     tagForm.State,
	}
	code = rpc.AddTag(tag)
	c.JSON(http.StatusOK, app.GetResponse(code, nil))
}

type EditTagForm struct {
	Id         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

func EditTag(c *gin.Context) {
	var tagForm EditTagForm
	httpCode, code := app.BindAndValid(c, &tagForm)
	if code != e.SUCCESS {
		c.JSON(httpCode, app.GetResponse(code, nil))
		return
	}
	tag := models.Tag{
		Model:      models.Model{ID: tagForm.Id,},
		Name:       tagForm.Name,
		CreatedBy:  "",
		ModifiedBy: tagForm.ModifiedBy,
		State:      tagForm.State,
	}
	code = rpc.EditTag(tag)
	c.JSON(http.StatusOK, app.GetResponse(code, nil))
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	vaild := validation.Validation{}
	vaild.Min(id, 1, "id").Message("ID不合法")

	if vaild.HasErrors() {
		app.MarkErrors(vaild.Errors)
		c.JSON(http.StatusOK, app.GetResponse(e.INVALID_PARAMS, nil))
		return
	}

	code := rpc.DeleteTag(id)
	c.JSON(http.StatusOK, app.GetResponse(code, nil))
}
