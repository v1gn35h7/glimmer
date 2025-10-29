# glimmer
Glimmer is a blazing-fast, Go-native stream processing engine that empowers developers to write distributed, stateful stream jobs using idiomatic Go.


## Usage
```
glimmer.NewStream("orders").
  Filter(func(e Event) bool { return e.Amount > 0 }).
  KeyBy(func(e Event) string { return e.UserID }).
  Window(TumblingWindow(10 * time.Second)).
  Aggregate(func(key string, events []Event) AggregatedEvent {
    total := 0.0
    for _, e := range events {
      total += e.Amount
    }
    return AggregatedEvent{UserID: key, Total: total}
  }).
  Sink(ToKafka("user-totals")).
  Run()

```
