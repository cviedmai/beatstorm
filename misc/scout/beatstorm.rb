require 'json'
require 'scout'

class BeatStormPlugin < Scout::Plugin
  def build_report
    return report(JSON.parse(File.read('/opt/beatstorm/stats_bs.json')))
  end
end
