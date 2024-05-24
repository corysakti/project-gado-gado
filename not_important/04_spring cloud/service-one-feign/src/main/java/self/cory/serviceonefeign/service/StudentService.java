package self.cory.serviceonefeign.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class StudentService {

    @Autowired
    private SubjectService subjectService;

    public String getStudentInfo() {
        return subjectService.getSubject();
    }
}
