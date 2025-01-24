package com.example.community_portal.demo.modal;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;
import java.util.UUID;

// Mails Entity
@Entity
@Table(name = "mails")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Mail {

    @Id
    @GeneratedValue
    private UUID id;

    @ManyToOne
    @JoinColumn(name = "author_club_id")
    private Club authorClub;

    @ManyToOne
    @JoinColumn(name = "author_user_id")
    private User authorUser;

    private String subject;

    private String content;

    private String recipients;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
}
