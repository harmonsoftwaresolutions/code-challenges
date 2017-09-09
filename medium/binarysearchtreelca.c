#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <strings.h>

struct Node {
	int data;
	struct Node *left, *right;
};

struct Node *newNode(int data) {
	struct Node *node = malloc(sizeof(struct Node));
	assert(node);
	node->data = data;
	node->left = node->right = NULL;

	return node;
};

struct Node *LCA(struct Node *root, int n1, int n2) {
	if (!root) return NULL;

	if (root->data > n1 && root->data > n2)
		return LCA(root->left, n1, n2);

	if (root->data < n1 && root->data < n2)
		return LCA(root->right, n1, n2);
	
	return root;
};

void PrintTree(struct Node *root) {
	printf("\t\t\t%d\n", root->data);
	printf("\t%d", root->left->data);
	printf("\t\t\t\t%d\n", root->right->data);
	printf("%d", root->left->left->data);
	printf("\t\t%d\n", root->left->right->data);

	return;
};

int main(int argc, char *argv[]) {
	struct Node *root = newNode(20);
	root->left = newNode(8);
	root->right = newNode(22);
	root->left->left = newNode(4);
	root->left->right = newNode(12);
	root->right->left = newNode(10);
	root->right->right = newNode(14);

	printf("\t\t *** BINARY TREE ***\n");
	PrintTree(root);
	printf("\n");

	int n1 = 10, n2 = 14;
	struct Node *t = LCA(root, n1, n2);
	printf("*** LCA Traverse\n");
	printf("LCA of %d and %d is %d\n", n1, n2, t->data);
	printf("\n");

	return 0;
};
