package self.cory.servicetwofeign;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;

@SpringBootApplication
@EnableFeignClients
public class ServiceTwoFeignApplication {

    public static void main(String[] args) {
        SpringApplication.run(ServiceTwoFeignApplication.class, args);
    }

}
