var app = angular.module("ath2014", ['ngRoute', 'ezfb', 'ui.codemirror','ngSanitize']);

app.config(function(ezfbProvider) {
  ezfbProvider.setInitParams({
    appId: '302908129857189'
  });
});

app.config(function($routeProvider, $locationProvider) {
  $routeProvider
    .when('/topic/create', {
      templateUrl: '/pages/add.html',
      controller: 'topicAddController'
    })
    .when('/topic/:permalink', {
      templateUrl: '/pages/topic.html',
      controller: 'topicController'
    })
    .when('/', {
      templateUrl: '/pages/home.html',
      controller: 'homeController'
    });
  $locationProvider.html5Mode(true);
});

var fbqlUrl = "https://api.facebook.com/method/fql.query?query="
var host = "ath2014.varokas.com"
function getFBLikeURL(permalink) {
    return fbqlUrl + encodeURI("select like_count from link_stat where url='" + host + "/topic/" + permalink) + "'&format=json"; 
}
function getScore(http, collector, topic) {
  http.get(getFBLikeURL(topic.Permalink)).success(function(data) {
    collector.push({
      "title": topic.Title,
      "permalink": topic.Permalink,
      "likes":  data[0]["like_count"]
    });

    collector.sort( function(a,b) { return b.likes - a.likes; } );
  });
}

app.controller("homeController", function($scope, $http, $routeParams) {
  $scope.topics = [];

  $http.get('/api/topics').success(function(data) {
     data.forEach( function(d) { getScore($http, $scope.topics, d); } );
  });
});

app.controller("topicController", function($scope, $location, $routeParams) {
  $scope.permalink = $routeParams.permalink;
  $scope.url = $location.absUrl();
});

app.controller("topicAddController", function($scope, $location, $http) {
  $scope.title = '';
  $scope.permalink = '';  
  $scope.preview = '';
  $scope.enableSubmit = true;

  $scope.submitForm = function(isValid) {
    if (isValid) {
      request = { "Title": $scope.title, "Permalink": $scope.permalink, "Description": $scope.description };
      $scope.enableSubmit = false;
      $http.post("/api/topics", request)
        .success(function() { $location.path('/'); })
        .error(function() { alert("error adding new topic"); $scope.enableSubmit=true; });
    }
  }

  $scope.generatePermalink = function() {
    if($scope.title != undefined && $scope.title.length > 0) {
      $scope.permalink = $scope.title.replace(/[^a-z0-9]+/gi, '-').replace(/^-*|-*$/g, '').toLowerCase();
    } else {
      $scope.permalink = "";
    }
  }

  $scope.generatePreview = function() {
    if($scope.description != undefined)  {
      $scope.preview = marked($scope.description);
    } else { 
      $scope.preview = ''
    }
  } 
});
