package com.example.community_portal.demo.modal;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;
import java.util.UUID;

// Comments Entity
@Entity
@Table(name = "comments")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Comment {

    @Id
    @GeneratedValue
    private UUID id;

    @ManyToOne
    @JoinColumn(name = "author_user_id")
    private User authorUser;

    @ManyToOne
    @JoinColumn(name = "post_id")
    private FeedPost post;

    @ManyToOne
    @JoinColumn(name = "comment_id")
    private Comment parentComment;

    private String comment;

    @Column(name = "like_count")
    private String likeCount;

    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @Column(name = "updated_at")
    private LocalDateTime updatedAt;
}
