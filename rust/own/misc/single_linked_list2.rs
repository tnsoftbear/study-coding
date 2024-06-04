struct Node<T> {
    value: T,
    next: Option<Box<Node<T>>>
}

struct List<T> {
    head: Option<Box<Node<T>>>
}

impl<T> List<T> {
    fn new() -> Self {
        Self {
            head: None
        }
    }

    fn push(&mut self, value: T) {
        let new_node = Node { 
            value,
            next: self.head.take()
        };
        self.head = Some(Box::new(new_node));
    }

    fn pop(&mut self) -> Option<T> {
        self.head.take().map(|node_box| {
            self.head = node_box.next;
            node_box.value
        })
    }

    fn peek(&self) -> Option<&T> {
        self.head.as_ref().map(|node_box| {
            &node_box.value
        })
    }

    pub fn peek_mut(&mut self) -> Option<&mut T> {
        self.head.as_mut().map(|node_box| {
            &mut node_box.value
        })
    }

    fn find_tail(&self) -> Option<&Box<Node<T>>> {
        self.head.as_ref().map(|mut node_box| {
            while node_box.next.is_some() {
                node_box = match node_box.next.as_ref() {
                    Some(n) => n,
                    None => panic!("unexpected")
                }
            }
            node_box
        })
    }

    fn find_tail_value(&self) -> Option<&T> {
        let mut current = self.head.as_ref();
        while let Some(node) = current {
            if node.next.is_none() {
                return Some(&node.value);
            }
            current = node.next.as_ref();
        }
        None
    }

}

////////////////////////////////////////////////////////////////////////////////

impl<T> Drop for List<T> {
    fn drop(&mut self) {
        let mut curr_opt = self.head.take();
        while let Some(mut node_box) = curr_opt {
            curr_opt = node_box.next.take();
        }
    }
}

////////////////////////////////////////////////////////////////////////////////

pub struct IntoIter<T>(List<T>);

impl<T> List<T> {
    pub fn into_iter(self) -> IntoIter<T> {
        IntoIter(self)
    }
}

impl<T> Iterator for IntoIter<T> {
    type Item = T;
    fn next(&mut self) -> Option<Self::Item> {
        self.0.pop()
    }
}

////////////////////////////////////////////////////////////////////////////////

impl<T> Iterator for List<T> {
    type Item = T;
    fn next(&mut self) -> Option<Self::Item> {
        self.pop()
    }
}

////////////////////////////////////////////////////////////////////////////////

pub struct Iter<'a, T> {
    next: Option<&'a Node<T>>
}

impl<T> List<T> {
    pub fn iter(&self) -> Iter<T> {
        Iter { 
            next: self.head.as_deref() 
        }
    }
}

impl<'a, T> Iterator for Iter<'a, T> {
    type Item = &'a T;
    fn next(&mut self) -> Option<Self::Item> {
        self.next.map(|node| {
            self.next = node.next.as_deref();
            &node.value
        })
    }
}

////////////////////////////////////////////////////////////////////////////////

pub struct IterMut<'a, T> {
    next: Option<&'a mut Node<T>>
}

impl<T> List<T> {
    pub fn iter_mut(&mut self) -> IterMut<T> {
        IterMut { 
            next: self.head.as_deref_mut() 
        }
    }
}

impl<'a, T> Iterator for IterMut<'a, T> {
    type Item = &'a mut T;
    fn next(&mut self) -> Option<Self::Item> {
        self.next.take().map(|node| {
            self.next = node.next.as_deref_mut();
            &mut node.value
        })
    }
}

////////////////////////////////////////////////////////////////////////////////

fn main() {
    // let n1 = Node { value: 10, next: None };
    // let n2 = Node { value: 20, next: Some(Box::new(n1)) };
    // let n3 = Node { value: 30, next: Some(Box::new(n2)) };
    // let n4 = Node { value: 40, next: None };
    // let list = LinkedList { head: Some(Box::new(n3)) };
}

#[cfg(test)]
mod test {
    use super::List;

    #[test]
    fn basics() {
        let mut list: List<i32> = List::new();
        list.push(10);
        list.push(20);
        assert_eq!(list.peek(), Some(&20));
        assert_eq!(list.peek_mut(), Some(&mut 20));
        list.peek_mut().map(|value| { *value = 22 });
        assert_eq!(list.pop(), Some(22));
        list.push(30);
        assert_eq!(list.find_tail().unwrap().value, 10);
        assert_eq!(list.pop(), Some(30));
        assert_eq!(list.pop(), Some(10));
        assert_eq!(list.pop(), None);
        assert!(matches!(list.find_tail(), None));
    }

    #[test]
    fn into_iter() {
        let mut list = List::new();
        list.push(1);
        list.push(2);
        list.push(3);
    
        let mut iter = list.into_iter();
        assert_eq!(iter.next(), Some(3));
        assert_eq!(iter.next(), Some(2));
        assert_eq!(iter.next(), Some(1));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn iterator_for_linked_list() {
        let mut list = List::new();
        list.push(1);
        list.push(2);
        list.push(3);
    
        let mut expected = 3;
        for item in list {
            assert_eq!(item, expected);
            expected -= 1;
        }
    }

    #[test]
    fn iter() {
        let mut list = List::new();
        list.push(1);
        list.push(2);
        list.push(3);
    
        let mut iter = list.iter();
        assert_eq!(iter.next(), Some(&3));
        assert_eq!(iter.next(), Some(&2));
        assert_eq!(iter.next(), Some(&1));
    }

    #[test]
    fn iter_mut() {
        let mut list = List::new();
        list.push(1);
        list.push(2);
        list.push(3);
    
        let mut iter = list.iter_mut();
        assert_eq!(iter.next(), Some(&mut 3));
        assert_eq!(iter.next(), Some(&mut 2));
        assert_eq!(iter.next(), Some(&mut 1));
    }
}
