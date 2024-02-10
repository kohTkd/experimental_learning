# frozen_string_literal: true

CONTAINERS = %w[api-main api-auth front].freeze
COMMANDS = %w[build run up stop restart reboot].freeze

File.open("#{__dir__}/../makefiles/docker.mk", 'w') do |f|
  CONTAINERS.each do |container|
    puts "Defining docker.#{container} commands..."
    f.puts ".PHONY: #{COMMANDS.map { |command| "docker.#{container}.#{command}" }.join(' ')}"
    COMMANDS.each do |command|
      puts "\tdocker.#{container}.#{command}"
      f.puts "docker.#{container}.#{command}:"
      f.puts "\tmake docker.#{command} container='#{container}'"
    end
    f.puts ''
    puts 'Done.'
  end
end
