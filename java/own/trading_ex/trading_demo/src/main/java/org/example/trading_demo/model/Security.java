package org.example.trading_demo.model;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Entity
@Table(name = "security")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class Security {
    @Id
    private long id;
    @Column(name = "name")
    private String name;
}
