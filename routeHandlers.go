package main

import(
	"math/rand"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func loginpage(c *gin.Context) {
    // Call the HTML method of the Context to render a template
    c.HTML(
        // Set the HTTP status to 200 (OK)
        http.StatusOK,
        // Use the index.html template
        "login.html",
        // Pass the data that the page uses (in this case, 'title')
        gin.H{
            "Title": "Login Page",
        },
    )  
}

func generateSessionToken() string {
    return strconv.FormatInt(rand.Int63(), 16)
}
 
func showregistrationPage(c *gin.Context) {    
   c.HTML(
       http.StatusOK,"register.html",
       gin.H{
        "Title": "Registration Page",
    }, 
    )
}

func register(c *gin.Context) {
     // Obtain the POSTed username and password values
    username := c.PostForm("username")
    password := c.PostForm("password")

    if _, err := registerNewUser(username, password); err == nil {
        // If the user is created, set the token in a cookie and log the user in
        token := generateSessionToken()
        c.SetCookie("gin_cookie", token, 3306, "/", "localhost", http.SameSiteLaxMode, false, true)
        c.Set("is_logged_in", true)
        c.HTML(http.StatusOK,
            "login_success.html" , 
            gin.H{
            "title": "Successful registration & Login"},
        )

    } else {
        // If the username/password combination is invalid,
        // show the error message on the login page
        c.HTML(http.StatusBadRequest, "register.html", gin.H{
            "ErrorTitle":   "Registration Failed",
            "ErrorMessage": err.Error()},
        )

    }
} 

func index(c *gin.Context){
    // Call the HTML method of the Context to render a template
    c.HTML(http.StatusOK,   
          "index.html",
          gin.H{
          "Title": "Home Page",
      },
    )
}


func login(c *gin.Context) {
  c.HTML(  
      http.StatusOK,
      "login_success.html",    
      gin.H{
          "Title": "Login Page",
      },
  )
}

func upload(c *gin.Context){
    c.HTML(  
        http.StatusOK,
        "upload.html",    
        gin.H{
            "Title": "upload Page",
        },
    )
}


/* func submitpic(w http.ResponseWriter, req *http.Request){
    in, header, err := req.FormFile("img")
    file,err := os.Open(C:\Users\Megha B Umesh\Pictures\Saved Pictures , handler.Filename)
} */

