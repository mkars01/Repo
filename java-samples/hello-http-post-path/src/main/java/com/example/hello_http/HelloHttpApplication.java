package com.example.hello_http;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
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

    @GetMapping("/{name}")
    public String sayHello(@PathVariable String name) {
        if (name == null || name.isEmpty()) {
            return "Hello, World (no Name parm)!";
        }
        return "Hello, " + name + "!";
    }
}
