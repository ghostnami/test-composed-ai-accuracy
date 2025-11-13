// src/components/UserProfile.jsx
import React from 'react';

// VULNERABILITY 8: XSS via dangerouslySetInnerHTML (Line 6)
export function UserProfile({ userBio }) {
  return <div dangerouslySetInnerHTML={{ __html: userBio }} />;
}

// VULNERABILITY 9: XSS via document.write (Line 10)
export function renderComment(comment) {
  document.write(`<div>${comment}</div>`);
}

export default UserProfile;

