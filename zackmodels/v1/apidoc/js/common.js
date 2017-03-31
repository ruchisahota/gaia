
const goToURLWithCurrentParameters = (url) => {
  window.location = url + getURLParameters();
}

const prepareComponents = (version) => {
  $('.navbar').sticky({
    context: 'body'
  });

  let parameters = getURLParameters();

  if (parameters.indexOf("fullscreen") > 0) {
    enterFullScreenMode();
  }
}

let getURLParameters = () => {
  return window.location.search;
}

let enterFullScreenMode = () => {
  $('.navbar').hide();
  $('#main_menu').hide();
  $('#main_content').attr("class", "sixteen wide column");
}
