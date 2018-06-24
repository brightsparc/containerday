#   Copyright 2017 Amazon.com, Inc. or its affiliates.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

from flask import Flask, jsonify
from flask.ext.cors import CORS, cross_origin
app = Flask(__name__)
cors = CORS(app)
app.config['CORS_HEADERS'] = 'Content-Type'

@app.route('/api')
@cross_origin()
def hello_world():
    return jsonify(response ="hi!  i\'m ALSO served via Python + Flask.  i\'m an API.")

if __name__ == '__main__':
    app.run(port=8000,host='0.0.0.0')
