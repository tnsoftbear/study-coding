import 'package:git_push/git_push.dart' as git_push;

void main(List<String> arguments) {
  if (arguments.length != 1) {
    print('Usage: dart script.dart <directory>');
    return;
  }

  String directory = arguments[0];
  git_push.commitAndPush(directory);
}
