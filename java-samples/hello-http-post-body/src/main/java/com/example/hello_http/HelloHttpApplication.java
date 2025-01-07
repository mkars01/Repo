package com.example.hello_http;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
public class HelloHttpApplication {

    public static void main(String[] args) {
        SpringApplication.run(HelloHttpApplication.class, args);
    }
}

@RestController
@RequestMapping("/hello")
class HelloController {

    // Handle POST requests with JSON body
    @PostMapping
    public String sayHello(@RequestBody NameRequest nameRequest) {
        if (nameRequest.getName() == null || nameRequest.getName().isEmpty()) {
            return "Hello, World (no Name parm)!";
        }
        return "Hello, " + nameRequest.getName() + "!";
    }
}

// Helper class to handle the JSON request body
class NameRequest {
    private String name;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
