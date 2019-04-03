var $ = jQuery.noConflict();
var formSubmitted = 'false';

jQuery(document).ready(function($) {	

	$('#formSuccessMessageWrap').hide(0);
	$('.formValidationError').fadeOut(0);
	
	// fields focus function starts
	$('input[type="text"], textarea').focus(function(){
		if($(this).val() == $(this).attr('data-dummy')){
			$(this).val('');
		};
	});
	// fields focus function ends
		
	// fields blur function starts
	$('input, textarea').blur(function(){
    	if($(this).val() == ''){
			$(this).val($(this).attr('data-dummy'));	
		};
	});
	// fields blur function ends
		
	// submit form data starts	   
    function submitData(currentForm, formType){     
		formSubmitted = 'true';		
		var formInput = $('#' + currentForm).serialize();		
		$.post($('#' + currentForm).attr('action'),formInput, function(data){			
			$('#' + currentForm).hide();
			$('#formSuccessMessageWrap').fadeIn(500);
		});

	};
	// submit form data function starts	
	// validate form function starts
	function validateForm(currentForm, formType){		
		// hide any error messages starts
	    $('.formValidationError').hide();
		$('.fieldHasError').removeClass('fieldHasError');
	    // hide any error messages ends	
		$('#' + currentForm + ' .requiredField').each(function(i){
			if($(this).val() == '' || $(this).val() == $(this).attr('data-dummy')){	
				$(this).val($(this).attr('data-dummy'));
				$(this).focus();
				$(this).addClass('fieldHasError');
				$('#' + $(this).attr('id') + 'Error').fadeIn(300);
				return false;
			};		
			if(formSubmitted == 'false' && i == $('#' + currentForm + ' .requiredField').length - 1){
			 	submitData(currentForm, formType);
			};  
   		});		
	};
	// validate form function ends	
	
	// contact button function starts
	$('#contentSubmitButton').click(function() {	
		validateForm($(this).attr('data-formId'));	
	    return false;
	});
	// contact button function ends
	
	
	
});
/*////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////*/
/*//////////////////// Document Ready Function Ends                                                                       */
/*////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////*/
