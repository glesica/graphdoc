# @Graph Blog
A data model for a blog, including posts, comments, users, and media files.

## @Node Post
A piece of content created by a user, can optionally include media files and
other users can comment on it.

### @Prop title

@Type str
@Index
@Required
@Unique

The post title, used to create a headline.

### @Prop content

@Type str

The content of the post, formatted as Markdown.

### @Prop html

@Type str

The content of the post, compiled from Markdown to HTML. This process happens
automatically whenever the post is edited.

### @Prop upvotes

@Type num

The number of votes the post has received from readers of the blog. This number
is incremented each time a user "upvotes" the post and decremented each time a
user "downvotes" the post.

## @Rel WRITTEN_BY

@From Post
@To User

A link from a post to the author or authors of the post. These are the people
who wrote the content and only they are allowed to make edits in the future.

## @Rel INCLUDES

@From Post
@To Media

A link from a post to a media item that is included in this post and can be used
by the authors within the post.

## @Node Comment
A comment on a post, created by a user (possibly even one of the authors).
A comment can only be associated with one post.

### @Prop content

@Type str
@Required

The content of the comment, formatted as Markdown.

### @Prop html

@Type str

The content of the comment,compiled from Markdown to HTML when the comment was
submitted. Comments cannot be edited, so this process only happens once.

## @Rel RELATED_TO

@From Comment
@To Post

A link from a comment to the post the comment was left on.

### @Rel WRITTEN_BY

@From Comment
@To User

A link from a comment to the author of the comment. There is only ever one
author.

## @Node User
A user of the blog, may optionally be allowed to create posts. All users are
allowed to create comments.

### @Prop post_author

@Type bool

Whether the user is permitted to create posts.

### @Prop name

@Type str

The screen name of the user, displayed on posts and comments.

### @Prop email

@Type str

The email address of the user, never displayed publicly but used internally.

## @Node Media
A media file available to be included in a post. Each file is owned by a
particular user and is only made available to that user when creating a post.

### @Prop path

@Type str

Path to the media file on the filesystem or a URL.

## @Rel OWNED_BY

@From Media
@To User

A link between a file and its owner. The file is only available to this user for
embedding in posts.
