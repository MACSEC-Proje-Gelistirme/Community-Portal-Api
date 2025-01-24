package com.example.community_portal.demo.modal;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;
import java.util.UUID;

// FeedPosts Entity
@Entity
@Table(name = "feed_posts")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class FeedPost {

    @Id
    @GeneratedValue
    private UUID id;

    @ManyToOne
    @JoinColumn(name = "author_club_id")
    private Club authorClub;

    @ManyToOne
    @JoinColumn(name = "author_user_id")
    private User authorUser;

    private String iamge;

    private String description;

    @Column(name = "like_count")
    private String likeCount;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
}
