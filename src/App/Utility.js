const openInNewTab = (url) => {
  window.open(url, '_blank');
};
const allValuesFalse = (obj) => Object.values(obj).every((i) => !i);

export { openInNewTab, allValuesFalse };
