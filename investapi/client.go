package investapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// DefaultEndpoint - это доменное имя по умолчанию, на котором работает InvestAPI
const DefaultEndpoint = "invest-public-api.tinkoff.ru"

// tokenAuth используется для описания стратегии авторизации GRPC клиента
type tokenAuth struct {
	// Token - это ключ доступа, который можно получить отсюда - // https://tinkoff.github.io/investAPI/grpc/#tinkoff-invest-api_1
	Token string
}

// GetRequestMetadata используется для авторизации с помощью Bearer стратегии
func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + t.Token,
	}, nil
}

// RequireTransportSecurity требует, чтобы транспорт был защищённым
func (tokenAuth) RequireTransportSecurity() bool {
	return true
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

// New создаёт новый клиент для доступа к API, используя Ключ доступа как аргумент
func New(token string) (client *Client, err error) {
	return NewWithCustomEndpoint(token, "invest-public-api.tinkoff.ru")
}

// NewWithCustomEndpoint cоздаёт новый клиент для доступа к API, используя Ключ доступа и доменное имя адреса API как аргументы
func NewWithCustomEndpoint(token, endpoint string) (client *Client, err error) {
	return NewWithOpts(token, endpoint, make([]grpc.DialOption, 0)...)
}

// NewWithOpts создаёт новый клиент для доступа к API, используя Ключ доступа как аргумент, доменное имя и вариадическую переменную opts с параметрами вида grpc.DialOption для настройки соединения
func NewWithOpts(token, endpoint string, opts ...grpc.DialOption) (client *Client, err error) {
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		ServerName: endpoint,
	})))
	opts = append(opts, grpc.WithPerRPCCredentials(tokenAuth{
		Token: token,
	}))
	conn, err := grpc.Dial(fmt.Sprintf("%s:443", endpoint), opts...)
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
