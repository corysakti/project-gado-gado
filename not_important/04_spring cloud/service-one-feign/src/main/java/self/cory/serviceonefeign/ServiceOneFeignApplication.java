package self.cory.serviceonefeign;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
public class ServiceOneFeignApplication {

    public static void main(String[] args) {
        SpringApplication.run(ServiceOneFeignApplication.class, args);
    }

}
