require 'json'
require 'scout'

class AuwfgPlugin < Scout::Plugin
  def build_report
    return report(JSON.parse(File.read('/opt/beatstorm/stats.json')))
  end
end
