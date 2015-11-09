(function(){
	// homeCtrl
	'use strict';

	console.log('Loading homeCtrl')

	angular.module('itrak').controller('homeCtrl', function($scope,$state,loginState){

		if (!loginState.loggedIn) {
			$state.go('login')
		}

		angular.extend($scope, {
			loginState: loginState,
		})
	});

})();

