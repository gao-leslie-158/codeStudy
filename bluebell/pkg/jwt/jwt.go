package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

// TokenExpireDuration token过期时间
const (
	aTokenExpireDuration = time.Hour * 1
	rTokenExpireDuration = time.Hour * 30
)

var (
	// 常见错误1：定义掩码，jwtSecret，这个key的底层是一个[]byte断言 --> key.([]byte)
	mySecret = []byte("春风不解风情")
)

// MyClaims 定义自己的数据
// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 怎么定义这个keyFunc并且使用
//var keyFunc jwt.Keyfunc
//type keyFunc func(*jwt.Token)(i interface{},err error){
//	return mySecret , err
//}

// GenToken 生成 Access Token 和 Refresh Token
func GenToken(UserID int64) (aToken, rToken string, err error) {
	// 创建一个自己的声明的数据
	c := MyClaims{
		UserID: UserID, // 自定义字段
		StandardClaims: jwt.StandardClaims{ // 标准字段
			ExpiresAt: time.Now().Add(
				time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "bluebell", // 签发人
		},
	}
	// 加密并获取完整编码后的token字符串
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)

	// refresh token 不需要存在任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(rTokenExpireDuration).Unix(),
		Issuer:    "bluebell",
	}).SignedString(mySecret)
	return
}

////GenToken 生成token
//func GenToken(UserID int64, Username string) (string, error) {
//	// 创建一个自己的声明的数据
//	c := MyClaims{
//		UserID:   UserID, // 自定义字段
//		Username: Username,
//		StandardClaims: jwt.StandardClaims{ // 标准字段
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
//			Issuer:    "bluebell",                                 // 签发人
//		},
//	}
//	// 使用指定的签名方法创建签名对象
//	// 常见错误2：jwt.NewWithClaims 他的加密方式应该选择jwt.SigningMethodHS256 而不是 jwt.SigningMethodES256 ，
//	// 这个H是hash的意思，而SigningMethodES256是没有SignedString方法的
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//	// 使用指定的secret签名并获取完整的编码后的字符串token
//	return token.SignedString(mySecret)
//}

// ParseToken 解析jwt , 解析Access Token
func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims) //解析结果变量，要初始化内存
	token, err = jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (i interface{}, err error) {
		return mySecret, err
	})
	if err != nil {
		return
	}
	// 对token对象中的Claims进行类型断言
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}

// RefreshToken 刷新 Access token
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, func(*jwt.Token) (i interface{}, err error) {
		return mySecret, err
	}); err != nil {
		return
	}

	// 从旧Access Token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(*jwt.Token) (i interface{}, err error) {
		return mySecret, err
	})
	v, _ := err.(*jwt.ValidationError)

	// 当Access token是过期错误，并且refresh没有过期就创建一个新的Access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}
