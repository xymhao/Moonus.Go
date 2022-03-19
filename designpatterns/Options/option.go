package Options

import (
	"time"
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeOut = 10
	defaultCaching = true
)

//Setting 将非必填的配置选项移到一个结构体里
type Setting struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(o *Setting)
}

// 定义一个函数类型，类似C#中的委托，委托参数为options
type optionFuc func(*Setting)

// Timeout 设置超时时间的函数
func Timeout(t time.Duration) optionFuc {
	return func(o *Setting) {
		o.timeout = t
	}
}

// Caching 设置缓存函数
func Caching(b bool) optionFuc {
	return func(o *Setting) {
		o.caching = b
	}
}

func InitServer(addr string, opts ...func(setting *Setting)) (*Setting, error) {
	server := Setting{}
	for _, opt := range opts {
		opt(&server)
	}
	return &server, nil
}

func (f optionFuc) apply(o *Setting) {
	f(o)
}

// WithTimeout 设置超时时间的函数
func WithTimeout(t time.Duration) Option {
	return optionFuc(func(o *Setting) {
		o.timeout = t
	})
}

// WithCaching 设置缓存函数
func WithCaching(b bool) Option {
	return optionFuc(func(o *Setting) {
		o.caching = b
	})
}

type ConnectBuilder struct {
	opts []Option
}

func (b ConnectBuilder) WithCache(cache bool) ConnectBuilder {
	b.opts = append(b.opts, WithCaching(cache))
	return b
}

func (b ConnectBuilder) WithTimeOut(t time.Duration) ConnectBuilder {
	b.opts = append(b.opts, WithTimeout(t))
	return b
}

func NewServer(addr string, opts ...Option) (*Connection, error) {
	options := Setting{
		timeout: defaultTimeOut,
		caching: defaultCaching,
	}

	for _, opt := range opts {
		opt.apply(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}

type User struct {
	Name  string
	Age   int
	Email string
	Phone string
	//...Other
	School string
	Gender string
}

// UserOption 我们定义一个函数类型，类似C#与委托函数Action
type UserOption func(*User)

func Phone(phone string) UserOption {
	return func(user *User) {
		user.Phone = phone
	}
}

func School(school string) UserOption {
	return func(user *User) {
		user.School = school
	}
}

func Gender(gender string) UserOption {
	return func(user *User) {
		user.Gender = gender
	}
}

func CreateUser(name string, age int, email string, opt ...UserOption) User {
	user := User{Name: name, Age: age, Email: email}

	for _, option := range opt {
		option(&user)
	}
	return user
}
