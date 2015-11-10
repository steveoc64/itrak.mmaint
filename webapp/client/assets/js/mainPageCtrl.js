(function(){
	// mainPageCtrl
	'use strict';

	console.log('Loading MainPage Controller')

	angular.module('itrak').controller('mainPageCtrl', function($state,loginState){

		angular.extend(this, {
			loginState: loginState,
			isLoggedIn: function() {
				return loginState.loggedIn
			},
			logout: function() {
				loginState.logout()
				$state.go("login")
			}
		})
	});

})();
