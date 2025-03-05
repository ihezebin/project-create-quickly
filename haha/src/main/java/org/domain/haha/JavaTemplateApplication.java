package org.domain.haha;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.boot.web.servlet.ServletComponentScan;
import org.springframework.scheduling.annotation.EnableScheduling;

import org.domain.haha.config.CustomizeConfig;

@SpringBootApplication
@ServletComponentScan
@EnableScheduling
@EnableConfigurationProperties(CustomizeConfig.class)
public class JavaTemplateApplication {

	public static void main(String[] args) {
		SpringApplication.run(JavaTemplateApplication.class, args);
	}

}
