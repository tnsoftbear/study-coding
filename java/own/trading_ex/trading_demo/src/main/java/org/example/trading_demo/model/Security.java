package org.example.trading_demo.model;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "security")
@Getter
@Setter
public class Security {
    @Id
    private long id;
    @Column(name = "name")
    private String name;
}
