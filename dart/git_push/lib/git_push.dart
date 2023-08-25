import 'dart:io';
import 'package:intl/intl.dart';

void commitAndPush(String directory) {
  Directory.current = Directory(directory);

  Process.run('git', ['add', '.'], runInShell: true).then((addResult) {
    if (addResult.exitCode == 0) {
      print('Git add successful.');
      Process.run('git', ['commit', '-am', getCurrentDateFormatted()],
              runInShell: true)
          .then((commitResult) {
        if (commitResult.exitCode == 0) {
          print('Commit successful.');
          Process.run('git', ['pull', '--rebase'], runInShell: true).then((pullResult) {
            if (pullResult.exitCode == 0) {
              print('Pull with rebase successful.');
              pushChanges();
            } else {
              print('Pull with rebase failed. Error: ${pullResult.stderr}');
            }
          });
        } else {
          print('Commit failed. Error: ${commitResult.stderr}');
        }
      }).catchError((error) {
        print('Git Commit caught error: $error');
      });
    } else {
      print('Git add failed. Error: ${addResult.stderr}');
    }
  }).catchError((error) {
    print('Git Add caught error: $error');
  });
}

void pushChanges() {
  Process.run('git', ['push'], runInShell: true).then((pushResult) {
    if (pushResult.exitCode == 0) {
      print('Push successful.');
    } else {
      print('Push failed. Error: ${pushResult.stderr}');
    }
  }).catchError((error) {
    print('Git Push caught error: $error');
  });
}

String getCurrentDateFormatted() {
  DateTime now = DateTime.now();
  String formattedDate = DateFormat('yyyy-MM-dd HH:mm').format(now);
  return formattedDate;
}
