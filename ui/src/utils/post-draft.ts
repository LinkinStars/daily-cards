const saveDraft = (content : string) => {
  localStorage.setItem("post-draft", content);
}

const readDraft = () => {
  return localStorage.getItem("post-draft");
}

const removeDraft = () => {
  localStorage.removeItem("post-draft");
}

export { saveDraft, readDraft, removeDraft };