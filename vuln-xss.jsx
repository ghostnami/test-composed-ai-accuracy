import React from 'react';

// VULNERABLE: Dangerously setting HTML
export function UserProfile({ userBio }) {
  return <div dangerouslySetInnerHTML={{ __html: userBio }} />;
}

// VULNERABLE: Direct innerHTML
export function CommentDisplay({ comment }) {
  const div = document.createElement('div');
  div.innerHTML = comment;
  return div;
}
