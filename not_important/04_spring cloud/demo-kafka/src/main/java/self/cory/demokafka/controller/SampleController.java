package self.cory.demokafka.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.Message;
import org.springframework.messaging.support.MessageBuilder;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;
import self.cory.demokafka.common.Sample;

@RestController
public class SampleController {
    @Autowired
    private KafkaTemplate<Object, Object> template;

    @PostMapping(path = "/send/{name}")
    public void sendFoo(@PathVariable String name) {
        this.template.send("sample-topic4", new Sample(name));
        Message<String> message = MessageBuilder
                .withPayload("stringValue")
                .setHeader(KafkaHeaders.TOPIC, "lokomotifdata")
                .build();
        template.send(message);
    }
}
