 $("#login-button").click(function(event){
	//event.preventDefault();

	 $('form').fadeOut(500);
	 $('.wrapper').addClass('form-success');

   var delayInMilliseconds = 1000; //1 second
   setTimeout(function() {
     $('form').submit();
   }, delayInMilliseconds);
});
