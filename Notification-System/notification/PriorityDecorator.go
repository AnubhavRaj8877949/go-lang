package notification

import "fmt"

type PriorityDecorator struct {
	INotificationDecorator
	Level string
}

func NewPriorityDecorator(n INotification, level string) *PriorityDecorator {
	return &PriorityDecorator{
		INotificationDecorator: INotificationDecorator{notification: n},
		Level:                  level,
	}
}

func (s *PriorityDecorator) GetContent() string {
	return fmt.Sprintf("%s\n-- %s", s.notification.GetContent(), s.Level)
}

// +-----------------------+          1..*          +-----------------------+
// |  RestaurantManager    |<>--------------------->|      Restaurant       |
// |     <<Singleton>>     |                        |       <<Model>>       |
// +-----------------------+                        +-----------------------+
// | - restaurants: List   |                        | - id: int             |
// +-----------------------+                        | - name: string        |
// | + AddRestaurant()     |                        | - address: string     |
// | + RemoveRestaurant()  |                        | - menu: List<MenuItem>|
// | + SearchByLocation()  |                        +-----------------------+
// +-----------------------+                        | + AddMenuItem()       |
//                                                  | + RemoveMenuItem()    |
//                                                  | + UpdateMenuItem()    |
//                                                  +-----------------------+
//                                                              | 1
//                                                              | (Composition: ◆)
//                                                              v 1..*
// +-----------------------+          1..*          +-----------------------+
// |         Cart          |<>--------------------->|       MenuItem        |
// +-----------------------+                        |       <<Model>>       |
// | - restaurant: Rest.   |                        +-----------------------+
// | - items: List<Item>   |                        | - code: int           |
// | - total: double       |                        | - name: string        |
// +-----------------------+                        | - price: double       |
// | + AddToCart(item)     |                        +-----------------------+
// | + Clear()             |
// | + IsEmpty(): bool     |
// +-----------------------+
//             ^
//             | (Composition: ◆)
// +-----------------------+                        +-----------------------+
// |         User          |----------------------->|         Order         |
// |       <<Model>>       |        places          |  <<Abstract Class>>   |
// +-----------------------+                        +-----------------------+
// | - id: int             |                        | # id: int             |
// | - name: string        |                        | # restaurant: Rest.   |
// | - address: string     |                        | # items: List<Item>   |
// | - cart: Cart          |                        | # user: User          |
// +-----------------------+                        | # payment: PayStrategy|
//                                                  +-----------------------+
//                                                  | + GetType(): string   |
//                                                  | + ProcessPayment()    |
//                                                  +-----------------------+
//                                                    ^         ^         |
//                                                    |         |         | uses
//                          (Inheritance: ───▷) ------+         |         v
//                          |                                   |   +-------------------+
//                +-------------------+               +-------------------+  PaymentStrategy  |
//                |   DeliveryOrder   |               |    PickupOrder    |  |   <<Interface>>   |
//                +-------------------+               +-------------------+  +-------------------+
//                | - deliveryAddress |               | - pickupTime      |  | + Pay(amount)     |
//                +-------------------+               +-------------------+  +-------------------+
//                | + GetType()       |               | + GetType()       |        ^       ^
//                +-------------------+               +-------------------+        |       |
//                          ^                                   ^        (Implements: - - -▷)
//                          |                                   |            +-----+       +-----+
//                          + - - - - - - - + - - - - - - - - - +            |                   |
//                                          | (Creates: - - ->)     +----------------+  +----------------+
//                                +-------------------+             |   UPIPayment   |  | CreditCardPay  |
//                                |   OrderFactory    |             +----------------+  +----------------+
//                                |   <<Interface>>   |             | + Pay(amount)  |  | + Pay(amount)  |
//                                +-------------------+             +----------------+  +----------------+
//                                | + CreateOrder()   |
//                                +-------------------+
//                                      ^       ^
//                                      |       | (Implements: - - -▷)
//                            +---------+       +---------+
//                            |                           |
//                  +-------------------+       +-------------------+
//                  |  NowOrderFactory  |       | ScheduleOrderFact.|
//                  +-------------------+       +-------------------+
//                  | + CreateOrder()   |       | + CreateOrder()   |
//                  +-------------------+       +-------------------+

// +-----------------------+          1..*          +-----------------------+
// |     OrderManager      |<>--------------------->|         Order         |
// |     <<Singleton>>     |                        +-----------------------+
// +-----------------------+
// | - orders: List<Order> |
// +-----------------------+
// | + AddOrder(order)     |
// | + GetAllOrders()      |
// +-----------------------+
