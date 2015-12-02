var child_process = require('child_process');

exports.handler = function(event, context) {
    var proc = child_process.spawn('./filter', [ JSON.stringify(event) ], {stdio: 'inherit'});

    proc.on('close', function(code) {
        if(code !== 0) {
            return context.done(new Error("process exited with non-zero status code"));
        }
        context.done(null);
    });
}
