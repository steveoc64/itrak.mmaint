(function(){
	// homeCtrl
	'use strict';

	console.log('Loading homeCtrl')

	angular.module('itrak').controller('homeCtrl', function($scope,loginState){

		if (!loginState.loggedIn) {
			$state.log('login')
		}
		
		angular.extend($scope, {
			loginState: loginState
		})
	});

})();

