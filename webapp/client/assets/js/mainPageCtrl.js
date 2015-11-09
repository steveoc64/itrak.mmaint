(function(){
	// mainPageCtrl
	'use strict';

	console.log('Loading mainPageCtrl')

	angular.module('itrak').controller('mainPageCtrl', function($scope,$state,loginState){

		angular.extend($scope, {
			loginState: loginState,
			isLoggedIn: function() {
				return loginState.loggedIn
			},
			logout: function() {
				loginState.logout()
				$state.go('/')
			}
		})
	});

})();
