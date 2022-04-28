const path = require('path');
  

module.exports = {
/**
 *||npx webpack serve config ||
entry: './resources/js/app.js',
output: {
    filename: 'app.js',
    path: __dirname + '/public',
},
*/
    resolve: {
        alias: {
            '@': path.resolve('resources/js'),
        },
	extensions:['.js','.jsx'],
    },


};
