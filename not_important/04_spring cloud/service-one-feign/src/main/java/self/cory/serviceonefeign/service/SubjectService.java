package self.cory.serviceonefeign.service;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.RequestMapping;

@FeignClient("subject")
public interface SubjectService {
    @RequestMapping("/api/subject")
    String getSubject();
}
