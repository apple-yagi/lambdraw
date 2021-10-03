import * as fs from 'fs';
import fetch from 'node-fetch';

fs.readFile(
  '/Users/ryuya/workspace/resize-api/tmp/original/gopher.png',
  async function (err, content) {
    if (err) {
      console.error(err);
    } else {
      const res = await fetch(
        'https://3d8r7a230b.execute-api.ap-northeast-1.amazonaws.com/default/resize-api',
        {
          method: 'POST',
          body: content
        }
      );
      console.log(await res.json());
    }
  }
);
