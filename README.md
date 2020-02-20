# serialbeat

[Benben's](https://github.com/benben/serialbeat) beat for shipping serial output to logstash or 
elasticsearch. But with an updated Beats version ([7.6.0](https://github.com/elastic/beats/releases/tag/v7.6.0)) and dependancies

## Synopsis

Create a configuration file with the name serialbeat.yml and add

    serialbeat:
      device: /dev/ttyACM0
      baud: 38400
      delimiter: "\n"
    output.console:
      pretty: true

And run `./serialbeat -c serialbeat.yml -e -d "*"` to see your serial output.

## Installation

    mkdir -p ${GOPATH}/src/github.com/itsties/serialbeat
    git clone git@github.com:itsties/serialbeat.git ${GOPATH}/src/github.com/itsties/serialbeat
    cd ${GOPATH}/src/github.com/itsties/serialbeat
    glide i
    make

Now you should have a `serialbeat` binary in the same directory.

## Configuration

See `serialbeat.reference.yml` for all possible configuration options.

## Development

This beat was 100% written on a Raspberry Pi 3.

Check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

## License

See LICENSE
