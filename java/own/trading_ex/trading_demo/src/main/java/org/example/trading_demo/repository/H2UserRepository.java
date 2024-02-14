package org.example.trading_demo.repository;

import org.example.trading_demo.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface H2UserRepository extends JpaRepository<User, Long> {
    User findByUsername(String username);
}
