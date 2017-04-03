
const goToURLWithCurrentParameters = (url) => {
  window.location = url + window.location.search;
}

const goToHash = (hash) => {
  $('.navbar .item').removeClass("active");

  if (!hash)
  {
    $('.navbar a.item:first-of-type').addClass("active");
  }
  else {
    $('.navbar a[href="'+hash+'"]').addClass("active");
  }
}

const prepareComponents = () => {
  $('pre').each(function(i, block) {
    hljs.highlightBlock(block);
  });

  goToHash(window.location.hash);

  if (window.location.search.indexOf("fullscreen") > 0) {
    enterFullScreenMode();
  }

  $('.navbar').sticky({
    context: 'div#main_container'
  });
}

let enterFullScreenMode = () => {
  $('.navbar').hide();
  $('#main_menu').hide();
  $('#main_content').attr("class", "sixteen wide column");
}
