var MyApp = angular.module('MyApp', []);

MyApp.controller('cUserList', function ($scope) {
  $scope.users = [
    {'name': 'Bob',
     'title': 'Head of Marketing'},
    {'name': 'Janet',
     'title': 'Product Manager'},
    {'name': 'Thanos',
     'title': 'Head of Productivity Theraputics'}
  ];
});
