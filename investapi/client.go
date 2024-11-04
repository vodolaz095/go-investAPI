package investapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
)

// DefaultEndpoint - это адрес, на котором работает продукционное окружение InvestAPI
const DefaultEndpoint = "invest-public-api.tinkoff.ru:443"

// SandboxEndpoint - это адрес, на котором работает тестовое окружение
const SandboxEndpoint = "sandbox-invest-public-api.tinkoff.ru:443"

// tokenAuth используется для описания стратегии авторизации GRPC клиента
type tokenAuth struct {
	// Token - это ключ доступа, который можно получить отсюда - // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1
	Token string
	// Secure - флаг, включающий требование шифрования соединения
	Secure bool
}

// GetRequestMetadata используется для авторизации с помощью Bearer стратегии
func (ta tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + ta.Token,
	}, nil
}

// RequireTransportSecurity требует, чтобы транспорт был защищённым
func (ta tokenAuth) RequireTransportSecurity() bool {
	return ta.Secure
}

// Client - это структура для доступа к API, в которой есть соединение и сгенерированные клиенты для всех сервисов
type Client struct {
	Connection               *grpc.ClientConn
	InstrumentsServiceClient InstrumentsServiceClient
	UsersServiceClient       UsersServiceClient
	MarketDataServiceClient  MarketDataServiceClient
	OperationsServiceClient  OperationsServiceClient
	OrdersServiceClient      OrdersServiceClient
	StopOrdersServiceClient  StopOrdersServiceClient
}

// Ping проверяет работоспособность соединения на основе получения состояния пула соединений
func (c *Client) Ping(_ context.Context) (err error) {
	// Полезная статья
	// https://grpc.github.io/grpc/core/md_doc_keepalive.html
	state := c.Connection.GetState()
	switch state {
	case connectivity.Ready, connectivity.Idle:
		return nil
	default:
		return fmt.Errorf("некорректное состояние grpc соединения: %s", state)
	}
}

// New создаёт новый клиент для доступа к API, используя Ключ доступа как аргумент
func New(token string, opts ...grpc.DialOption) (client *Client, err error) {
	return NewWithCustomEndpoint(token, DefaultEndpoint, opts...)
}

// NewWithCustomEndpoint создаёт новый клиент для доступа к API, используя Ключ доступа и д оменное имя адреса API как аргументы
func NewWithCustomEndpoint(token, endpoint string, opts ...grpc.DialOption) (client *Client, err error) {
	return NewWithOpts(token, endpoint, opts...)
}

// NewWithOpts создаёт новый клиент для доступа к API, используя Ключ доступа как аргумент, доменное имя и вариадическую
// переменную opts с параметрами вида grpc.DialOption для настройки соединения
func NewWithOpts(token, endpoint string, opts ...grpc.DialOption) (client *Client, err error) {
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: strings.Split(endpoint, ":")[0],
		})),
		grpc.WithPerRPCCredentials(tokenAuth{Token: token, Secure: true}),
		grpc.WithUserAgent("https://github.com/vodolaz095/go-investAPI"),
	)
	return makeClient(endpoint, opts...)
}

// NewInsecure создаёт новый клиент с ВЫКЛЮЧЕННОЙ проверкой подлинности TLS сертификатов для доступа к API, используя
// Ключ доступа как аргумент, доменное имя и вариадическую переменную opts с параметрами вида grpc.DialOption для настройки соединения
// Этот клиент лучше не использовать за пределами тестовых окружений!
func NewInsecure(token, endpoint string, opts ...grpc.DialOption) (client *Client, err error) {
	opts = append(opts,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:         strings.Split(endpoint, ":")[0],
			InsecureSkipVerify: true, // ВАЖНО
		})),
		grpc.WithPerRPCCredentials(tokenAuth{Token: token, Secure: false}),
		grpc.WithUserAgent("https://github.com/vodolaz095/go-investAPI"),
	)
	return makeClient(endpoint, opts...)
}

func makeClient(endpoint string, opts ...grpc.DialOption) (client *Client, err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return
	}
	client = new(Client)
	client.Connection = conn
	client.InstrumentsServiceClient = NewInstrumentsServiceClient(conn)
	client.UsersServiceClient = NewUsersServiceClient(conn)
	client.MarketDataServiceClient = NewMarketDataServiceClient(conn)
	client.OperationsServiceClient = NewOperationsServiceClient(conn)
	client.OrdersServiceClient = NewOrdersServiceClient(conn)
	client.StopOrdersServiceClient = NewStopOrdersServiceClient(conn)
	return
}
