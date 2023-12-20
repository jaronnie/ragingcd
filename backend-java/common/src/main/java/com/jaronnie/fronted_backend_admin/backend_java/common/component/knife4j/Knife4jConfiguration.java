package com.jaronnie.fronted_backend_admin.backend_java.common.component.knife4j;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import springfox.documentation.builders.ApiInfoBuilder;
import springfox.documentation.builders.PathSelectors;
import springfox.documentation.builders.RequestHandlerSelectors;
import springfox.documentation.service.Contact;
import springfox.documentation.spi.DocumentationType;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2WebMvc;

@Configuration
@EnableSwagger2WebMvc
public class Knife4jConfiguration {

    @Bean
    public Docket webApiConfig() {
        return new Docket(DocumentationType.SWAGGER_2)
                .apiInfo(new ApiInfoBuilder()
                        .title("接口文档")
                        .description("backend-java")
                        .contact(new Contact("Jaronnie", "https://fronted-backend-admin.cloud.jaronnie.com", "jaron@jaronnie.com"))
                        .version("v1.0")
                        .build())
                .select()
                .apis(RequestHandlerSelectors.basePackage("com.jaronnie.fronted_backend_admin.backend_java"))
                .paths(PathSelectors.any())
                .build();
    }
}