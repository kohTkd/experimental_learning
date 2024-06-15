# frozen_string_literal: true

APP_CONTAINERS = %w[api-main api-auth front].freeze
APP_COMMANDS = %w[build run up stop restart reboot logs attach down].freeze

DB_CONTAINERS = %w[database-main database-auth].freeze
DB_COMMANDS = %w[build up stop logs down attach].freeze

File.open("#{__dir__}/../makefiles/docker.mk", 'w') do |f|
  [[APP_CONTAINERS, APP_COMMANDS], [DB_CONTAINERS, DB_COMMANDS]].each do |containers, commands|
    containers.each do |container|
      puts "Defining docker.#{container} commands..."
      f.puts ".PHONY: #{commands.map { |command| "docker.#{container}.#{command}" }.join(' ')}"
      commands.each do |command|
        puts "\tdocker.#{container}.#{command}"
        f.puts "docker.#{container}.#{command}:"
        f.puts "\tmake docker.#{command} container='#{container}'"
      end
      f.puts ''
      puts 'Done.'
    end
  end
end
