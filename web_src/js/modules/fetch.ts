// hunk: @@ -1,11 +1,6 @@
import {isObject} from '../utils.ts';
import type {RequestOpts} from '../types.ts';

// fetch wrapper, use below method name functions and the `data` option to pass in data
// which will automatically set an appropriate headers. For json content, only object
// and array types are currently supported.
// hunk: @@ -20,7 +15,6 @@ export function request(url: string, {method = 'GET', data, headers = {}, ...oth
  }

  const headersMerged = new Headers({
    ...(contentType && {'content-type': contentType}),
  });

