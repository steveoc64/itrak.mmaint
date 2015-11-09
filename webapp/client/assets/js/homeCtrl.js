(function(){
	// homeCtrl
	'use strict';

	console.log('Loading homeCtrl')

	angular.module('itrak').controller('homeCtrl', function($state,loginState){

		if (!loginState.loggedIn) {
			$state.go('login')
		}

		angular.extend(this, {
			loginState: loginState,
		})
	});

})();

