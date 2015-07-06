package controllers

// controllers pkg is for each action requests
// create a action func, you need to create a *.html as well.
import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// create a APP func here, you need to create a view hello.html
// in views/App/Hello.html
func (c App) Hello(myname string) revel.Result {
	return c.Render(myname)
}
